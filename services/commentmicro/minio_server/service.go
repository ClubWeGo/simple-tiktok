package minio_server

import (
	"bytes"
	"context"
	"fmt"

	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func UploadFile(objectName, filePath, contentType, bucketName string) (url string, err error) {
	// Upload the file with FPutObject
	_, err = MinioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}

	// log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	fileurl := GlobalConfig.Endpoint + "/" + bucketName + "/" + objectName
	return fileurl, nil
}

func UploadMP4VideoToDouYin(objectName, videoPath string, frameNum int32) (videourl, coverurl string, err error) {
	// Upload the mp4 file with FPutObject
	_, err = MinioClient.FPutObject(context.Background(), "douyin", objectName, videoPath, minio.PutObjectOptions{ContentType: "video/mp4"})
	if err != nil {
		return "", "", err
	}

	coverpicName := objectName + ".jpg"
	coverpath := fmt.Sprintf("/tmp/%s", coverpicName)

	// 使用ffmpeg 提取指定帧作为图像文件
	// 需要先安装ffmpeg, linux下 sudo apt install ffmpeg
	// 帧处理参考 https://juejin.cn/post/7198400665696813115
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).WithOutput(buf).Run()
	if err != nil {
		return "", "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		return "", "", err
	}

	err = imaging.Save(img, coverpath)
	if err != nil {
		return "", "", err
	}

	// TODO : 视频上传成功，封面操作失败，是否需要删掉视频？
	_, err = MinioClient.FPutObject(context.Background(), "douyin", coverpicName, coverpath, minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		return "", "", err
	}

	// log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	videourl = "http://" + GlobalConfig.Endpoint + "/douyin/" + objectName
	coverurl = "http://" + GlobalConfig.Endpoint + "/douyin/" + coverpicName
	return videourl, coverurl, nil
}
