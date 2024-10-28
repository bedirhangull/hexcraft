package cmd

import (
	"log"
	"os"

	"github.com/bedirhangull/hexcraft/pkg/config"
	"github.com/bedirhangull/hexcraft/pkg/logger"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var hexcraft = &cobra.Command{
	Use:   "generate",
	Short: "Generate a hexcraft project",
	Run:   generateFolderStructure,
}

func init() {
	rootCmd.AddCommand(hexcraft)
	hexcraft.Flags().StringP("path", "p", "", "Path to the config file")
}

func generateFolderStructure(cmd *cobra.Command, args []string) {
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		log.Fatal("Error getting path:", err)
	}
	if path == "" {
		pathInfo := logger.NewLogger("info", "Default path is used")
		pathInfo.Log()
		path = "hexcraft.yaml"
	}

	config, err := getConfig(path)
	if err != nil {
		log.Fatal("Error getting config:", err)
	}

	err = createFolderStructure(config)
	if err != nil {
		log.Fatal("Error creating folder structure:", err)
	}

	successLog := logger.NewLogger("success", "Successfully generated hexagonal architecture folder structure")
	successLog.Log()
}

func getConfig(path string) (*config.Config, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		yamlFileLog := logger.NewLogger("error", "Error reading the config file")
		yamlFileLog.Log()
		return nil, err
	}
	var cfg config.Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func createFolderStructure(config *config.Config) error {
	err := os.Mkdir(config.PackageName, os.ModePerm)
	if err != nil {
		return err
	}

	folders := []string{
		"cmd",
		"internal/adapter/config",
		"internal/adapter/grpcserver",
		"internal/adapter/proto",
		"internal/adapter/storage",
		"internal/core/domain",
		"internal/core/port",
		"internal/core/service",
		"pkg",
		"test",
	}

	for _, folder := range folders {
		err := os.MkdirAll(config.PackageName+"/"+folder, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if config.Dockerfile {
		err := createFile(config.PackageName + "/Dockerfile")
		if err != nil {
			return err
		}
	}

	if config.GitIgnore {
		err := createFile(config.PackageName + "/.gitignore")
		if err != nil {
			return err
		}
	}

	if config.EnvFile {
		err := createFile(config.PackageName + "/.env")
		if err != nil {
			return err
		}
	}

	return nil
}

func createFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
