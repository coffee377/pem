#! /bin/sh

#$GOARCH 目标平台（编译后的目标平台）的处理器架构（386,amd64,arm）
#$GOOS 目标平台（编译后的目标平台）的操作系统（darwin,freebsd,linux,windows）

#操作系统
OS=("linux" "windows" "darwin")
#处理器
ARCH=("386" "amd64" "arm")
BUILD_NAME=
ZIP_FILES=
#
for os in ${OS[@]}
do
	for arch in ${ARCH[@]}
	do
		if [[  $os != 'linux' && $arch == 'arm' ]]
		then
			continue
		fi
		# windows 系统追加后缀
		if [[ $os == 'windows' ]] ;then
			ext=".exe"
		else
			ext=""	
		fi
		BUILD_NAME="lic_${os}_${arch}${ext}"
		echo ${BUILD_NAME}
		ZIP_FILES="${ZIP_FILES} ./${BUILD_NAME}"
		echo ${ZIP_FILES}
		#CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -o ${BUILD_NAME}
	done
done

# 目录不存在则创建目录
if [ ! -d "./dist" ]; then
  mkdir -p -m 755 ./dist
fi

# 打包压缩文件
zip ./dist/lic.zip ${ZIP_FILES}


