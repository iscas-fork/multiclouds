package code_generator

import (
	cloud_manager "cloud_manager/src/analyzer"
	"cloud_manager/src/code_generator"
	"cloud_manager/src/codegen/openstack"
	"cloud_manager/src/utils"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"os"
	"path/filepath"
	"testing"
)

func TestGenAliyunRequestRegistry(t *testing.T) {
	client := ecs.Client{}
	analyzer := cloud_manager.NewCloudAPIAnalyzer()
	analyzer.ExtractCloudAPIs(client)
	requestRegistryInfo := analyzer.ExtractRequestInfos("Create")

	templatePath := "E:\\gopath\\src\\cloud_manager\\src\\code_generator\\templates\\request_map.tmpl"
	requestRegistryInfo.ImportPaths = []string{"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"}
	data := utils.Struct2Map(requestRegistryInfo)
	params := map[string]interface{}{
		"packageName": "registry",
		"kind":        "Aliyun",
	}
	code, err := code_generator.GenCode(templatePath, data, params)
	if err != nil {
		t.Error(err)
	}
	filePtr, err := os.Create("E:\\gopath\\src\\cloud_manager\\src\\codegen\\registry\\aliyun_create_request_registry.go")
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	defer filePtr.Close()
	filePtr.Write(code)
}
func TestGenOpenstackRequestRegistry(t *testing.T) {
	client := openstack.OpenstackClient{}
	analyzer := cloud_manager.NewCloudAPIAnalyzer()
	analyzer.ExtractCloudAPIs(client)
	requestRegistryInfo := analyzer.ExtractRequestInfos("New")

	templatePath := "E:\\gopath\\src\\cloud_manager\\src\\code_generator\\templates\\request_map.tmpl"
	requestRegistryInfo.ImportPaths = []string{"cloud_manager/src/codegen/openstack"}
	data := utils.Struct2Map(requestRegistryInfo)
	params := map[string]interface{}{
		"packageName": "registry",
		"kind":        "Openstack",
	}
	code, err := code_generator.GenCode(templatePath, data, params)
	if err != nil {
		t.Error(err)
	}
	filePtr, err := os.Create("E:\\gopath\\src\\cloud_manager\\src\\codegen\\registry\\openstack_create_request_registry.go")
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	defer filePtr.Close()
	filePtr.Write(code)
}
func TestGenOpenstackCode(t *testing.T) {
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\compute\\v2\\extensions\\secgroups"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\objectstorage\\v1\\containers"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\blockstorage\\v3\\qos"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\containerinfra\\v1\\clusters"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\compute\\v2\\extensions\\bootfromvolume"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\compute"
	dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\compute\\v2\\servers"
	ma := cloud_manager.NewModuleAnalyzer()
	resourceInfos, err := ma.DoAnalyze(dir)
	if err != nil {
		t.Error(err)
	}
	for _, resourceInfo := range resourceInfos {
		if len(resourceInfo.ActionInfos) == 0 {
			continue
		}
		fmt.Printf("gen code for actions in resource %s\n", resourceInfo.ResourcePackageName)
		templatePath := "E:\\gopath\\src\\cloud_manager\\src\\code_generator\\templates\\openstack_request.tmpl"
		data := utils.Struct2Map(resourceInfo)
		params := map[string]interface{}{
			"packageName": "openstack",
		}
		code, err := code_generator.GenCode(templatePath, data, params)
		if err != nil {
			t.Error(err)
		}
		basePath := "E:\\gopath\\src\\cloud_manager\\src\\codegen\\openstack"
		fileName := utils.JoinName(resourceInfo.ResourcePath, "openstack", "_") + "_request.go"
		filePath := filepath.Join(basePath, fileName)
		filePtr, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			t.Error(err)
		}
		defer filePtr.Close()
		filePtr.Write(code)
	}
}

func TestGenOpenstackResultCode(t *testing.T) {
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\compute\\v2\\extensions\\secgroups"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\objectstorage\\v1\\containers"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\blockstorage\\v3\\qos"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\containerinfra\\v1\\clusters"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\compute\\v2\\servers"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\compute\\v2\\extensions\\bootfromvolume"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\compute\\v2\\extensions\\quotasets"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\compute\\v2\\extensions\\quotasets"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\clustering\\v1\\profiletypes"
	//dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack\\container\\v1\\capsules"
	dir := "E:\\gopath\\pkg\\mod\\github.com\\gophercloud\\gophercloud@v1.0.0\\openstack"
	ma := cloud_manager.NewModuleAnalyzer()
	resourceInfos, err := ma.DoAnalyze(dir)
	if err != nil {
		t.Error(err)
	}
	for _, resourceInfo := range resourceInfos {
		if len(resourceInfo.ActionInfos) == 0 || !resourceInfo.HasResultFile {
			continue
		}
		fmt.Printf("gen code for actions in resource %s\n", resourceInfo.ResourcePackageName)
		templatePath := "E:\\gopath\\src\\cloud_manager\\src\\code_generator\\templates\\openstack_result.tmpl"
		data := utils.Struct2Map(resourceInfo)
		params := map[string]interface{}{
			"packageName": "openstack",
		}
		code, err := code_generator.GenCode(templatePath, data, params)
		if err != nil {
			t.Error(err)
		}
		basePath := "E:\\gopath\\src\\cloud_manager\\src\\codegen\\openstack"
		fileName := utils.JoinName(resourceInfo.ResourcePath, "openstack", "_") + "_result.go"
		filePath := filepath.Join(basePath, fileName)
		filePtr, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			t.Error(err)
		}
		defer filePtr.Close()
		filePtr.Write(code)
	}
}
