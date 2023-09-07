package main

import (
	tools "fbx-to-tileset/pkg"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/alecthomas/kingpin"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	// logrus.SetReportCaller(true)
}

func CheckBinExists(paths []string) error {
	// todo
	return nil
}

func ExecuteCommand(name string, args []string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Fbx2Tileset(input, output, texture string, clear bool, lng, lat, height float64, verticesLimit, minSize uint32) {
	fbxFolder := "./" + tools.UUID() + "/"
	gltfFolder := "./" + tools.UUID() + "/"

	// fbx-split
	if err := ExecuteCommand("./fbx-split", []string{"-i", input, "-o", fbxFolder}); err != nil {
		log.Panic(err)
	}

	// texture
	if texture != "" {
		log.Printf("Copy texture from %s to %s\n", texture, fbxFolder)
		if err := ExecuteCommand("/bin/bash", []string{"-c", fmt.Sprintf("cp %s %s", texture+"*", fbxFolder)}); err != nil {
			log.Panic(err)
		}
	}

	// FBX2glTF
	if err := ExecuteCommand("./FBX2glTF.sh", []string{fbxFolder, gltfFolder}); err != nil {
		log.Panic(err)
	}

	// gltf-to-tileset
	if err := ExecuteCommand("./gltf-to-tileset", []string{
		"-i", gltfFolder,
		"-o", output,
		"--lng", strconv.FormatFloat(lng, 'f', 10, 64),
		"--lat", strconv.FormatFloat(lat, 'f', 10, 64),
		"--height", strconv.FormatFloat(height, 'f', 10, 64),
		"-v", strconv.FormatUint(uint64(verticesLimit), 10),
		"-m", strconv.FormatUint(uint64(minSize), 10)}); err != nil {
		log.Panic(err)
	}

	// clear
	if clear {
		// todo remove fbxfloder & glbfolder
		log.Printf("Remove folder %s\n", fbxFolder+" "+gltfFolder)
		if err := ExecuteCommand("rm", []string{"-rf", fbxFolder, gltfFolder}); err != nil {
			log.Panic(err)
		}
	}
}

func main() {
	var (
		input         = kingpin.Flag("input", "输入的FBX模型文件").Short('i').Required().String()
		output        = kingpin.Flag("output", "输出切片文件的目录.").Default("./out/").Short('o').String()
		texture       = kingpin.Flag("texture", "贴图文件目录，如果设置了则会进行尺寸优化。如果没有设置，则依据FBX2glTF的规则进行贴图搜寻，且不对贴图做任何处理，搜寻规则为FBX文件目录/.fbm目录/当前工作目录").Short('t').String()
		clear         = kingpin.Flag("clear", "清理临时目录").Short('c').Bool()
		lng           = kingpin.Flag("lng", "生成切片的经度").Default("39.90691").Float64()
		lat           = kingpin.Flag("lat", "生成切片的纬度").Default("116.39123").Float64()
		height        = kingpin.Flag("height", "生成切片的高度").Default("0").Float64()
		verticesLimit = kingpin.Flag("verticeslimit", "单个b3dm文件的最大顶点数").Short('v').Default("500000").Uint32()
		minSize       = kingpin.Flag("minsize", "单个模型文件的最小尺寸，小于该尺寸的会合并成cmpt文件").Short('m').Default("2").Uint32()
	)

	kingpin.Parse()

	Fbx2Tileset(*input, tools.FixFolderPath(*output), tools.FixFolderPath(*texture), *clear, *lng, *lat, *height, *verticesLimit, *minSize)
}
