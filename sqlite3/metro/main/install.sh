echo "install start"
go build
cp main.exe /D/workspace/projects/golang/src/evolve/evolution/projects/metro/reportServer/bin/report.exe
cp ./template/* /D/workspace/projects/golang/src/evolve/evolution/projects/metro/reportServer/bin/template
echo "install finish"