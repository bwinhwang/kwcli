package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var complianceDownloadCmd = &cobra.Command{
	Use:   "compliance_download",
	Short: "Download compliance report file.",
	Run: func(cmd *cobra.Command, args []string) {
		// Collect parameters
		project, _ := cmd.Flags().GetString("project")
		filePath, _ := cmd.Flags().GetString("file_path")

		// Get auth info directly
		url, user, token, err := LoadKWauthInfo()
		if err != nil {
			fmt.Printf("Error loading auth info: %v\n", err)
			return
		}

		// Construct the correct URL for file download
		// Based on the HTML response, it seems we need to use a different endpoint
		baseURL := strings.TrimSuffix(url, "/review/api")
		downloadURL := fmt.Sprintf("%s/review/api?action=compliance_download&user=%s&ltoken=%s&project=%s&file_path=%s",
			baseURL,
			user,
			token,
			project,
			filePath)

		// Make HTTP POST request to download file (Klocwork API typically uses POST)
		req, err := http.NewRequest(http.MethodPost, downloadURL, nil)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			return
		}

		// Set appropriate headers
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept", "application/octet-stream")

		// Create HTTP client and make the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error downloading file: %v\n", err)
			return
		}
		defer resp.Body.Close()

		// Check if response is successful
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error downloading file: HTTP %d\n", resp.StatusCode)
			// Read response body for more details
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Response: %s\n", string(body))
			return
		}

		// Check content type to ensure it's not HTML
		contentType := resp.Header.Get("Content-Type")
		if strings.Contains(contentType, "text/html") || strings.Contains(contentType, "application/xhtml+xml") {
			fmt.Printf("Error: Received HTML response instead of file content. This indicates authentication or endpoint issues.\n")
			fmt.Printf("Content-Type: %s\n", contentType)
			return
		}

		// Extract filename from file_path or use content-disposition header
		filename := filepath.Base(filePath)
		if filename == "" || filename == "." || filename == "/" {
			filename = "compliance_report"
		}

		// Check if content-disposition header provides a filename
		// if contentDisposition := resp.Header.Get("Content-Disposition"); contentDisposition != "" {
		// 	if start := strings.Index(contentDisposition, "filename="); start != -1 {
		// 		filename = contentDisposition[start+9:]
		// 		if end := strings.Index(filename, ";"); end != -1 {
		// 			filename = filename[:end]
		// 		}
		// 		filename = strings.Trim(filename, `"`)
		// 	}
		// }

		// Create output file
		outFile, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return
		}
		defer outFile.Close()

		// Copy downloaded content to file
		_, err = io.Copy(outFile, resp.Body)
		if err != nil {
			fmt.Printf("Error saving file: %v\n", err)
			return
		}

		fmt.Printf("File downloaded successfully: %s\n", filename)
	},
}

func init() {
	// Required flags
	complianceDownloadCmd.Flags().StringP("project", "p", "", "Project name (required)")
	complianceDownloadCmd.MarkFlagRequired("project")
	complianceDownloadCmd.Flags().StringP("file_path", "f", "", "Relative path to the report file (returned by the compliance report generate method) (required)")
	complianceDownloadCmd.MarkFlagRequired("file_path")

	// Record flags in definedFlags
	definedFlags["project"] = true
	definedFlags["file_path"] = true

	rootCmd.AddCommand(complianceDownloadCmd)
}
