package file

import (
	"bufio"
	"image"
	"image/color"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

const (
	PNG  = "png"
	JPG  = "jpg"
	JPEG = "jpeg"
	TIF  = "tif"
	GIF  = "gif"
)

func imageTypeByExt(ext string) string {
	switch ext {
	case ".jpg":
		return JPEG
	case ".png":
		return PNG
	case ".tif":
		return TIF
	}
	return JPEG
}

func GetImage(w io.Writer, imgPath, prdtType, subType string, width, height int, fit string) (imgType string, err error) {
	// 原图像
	imageFile, err := os.Open(imgPath)
	if err != nil {
		return JPEG, err
	}
	defer imageFile.Close()

	ext := filepath.Ext(imgPath)
	imgType = imageTypeByExt(ext)

	// 没有宽高，就是在加载原图像
	if width == 0 && height == 0 {
		io.Copy(w, imageFile)
		return imgType, nil
	}

	// 裁剪图像的组合路径
	if fit == "" {
		fit = "thumbnail"
	}

	filename := filepath.Base(imgPath)

	fileWithoutExt := strings.Trim(filename, ext)
	var str = strings.Builder{}
	str.WriteString(fileWithoutExt)
	str.WriteString("_")
	str.WriteString(strconv.Itoa(width))
	str.WriteString("_")
	str.WriteString(strconv.Itoa(height))
	str.WriteString("_" + fit)
	str.WriteString(ext)
	cutPathDir := filepath.Dir(imgPath)
	cutPath := filepath.Join(cutPathDir, str.String())

	// 判断是否存在裁剪图像
	if cutFile, err := os.Open(cutPath); err == nil {
		io.Copy(w, cutFile)
		cutFile.Close()
		return imgType, nil
	}

	// 图片解码 --------------------------------------
	bufFile := bufio.NewReader(imageFile)
	srcImg, imgType2, err := image.Decode(bufFile)
	if err != nil {
		return imgType2, err
	}

	// 要裁剪的宽高不能大于自身的宽高
	Rwidth := srcImg.Bounds().Max.X
	if width > Rwidth {
		width = Rwidth
	}

	Rheight := srcImg.Bounds().Max.Y
	if height > Rheight {
		height = Rheight
	}

	// gif 图就不处理了
	if imgType2 == GIF || (width == Rwidth && height == Rheight) {
		// 设置文件的偏移量 - 因为文件被 image.Decode 后文件的偏移量到尾部
		imageFile.Seek(0, 0)
		// 向浏览器输出
		io.Copy(w, imageFile)
		return imgType2, err
	}

	var dstImage *image.NRGBA
	switch fit {
	case "fill":
		dstImage = imaging.Fill(srcImg, width, height, imaging.Center, imaging.Lanczos)
	case "fit":
		dstImage = imaging.Fit(srcImg, width, height, imaging.Lanczos)
	case "crop":
		dstImage = imaging.Crop(srcImg, image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{width, height}})
	case "resize":
		dstImage = imaging.Resize(srcImg, width, height, imaging.Lanczos)
	case "thumbnail":
		dstImage = imaging.Thumbnail(srcImg, width, height, imaging.Lanczos)
	default:
		dstImage = imaging.Thumbnail(srcImg, width, height, imaging.Lanczos)
	}

	dst := imaging.New(width, height, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, dstImage, image.Pt(0, 0))

	imaging.Save(dst, cutPath)
	os.Chmod(cutPath, 0777)

	if imgType == JPEG || imgType == JPG {
		jpeg.Encode(w, dst, nil)
	} else if imgType == PNG {
		png.Encode(w, dst)
	}

	return imgType, nil
}
