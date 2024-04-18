# Function to build for a specific OS/architecture
build_for_os_arch() {
    os=$1
    arch=$2
    output_dir="../bin/${os}-${arch}"

    echo "Building for ${os}/${arch}"

    # Ensure the output directory exists
    mkdir -p "${output_dir}"

    # Build the kwcli executable with appropriate GOOS and GOARCH
    env CGO_ENABLED=0  GOOS="${os}" GOARCH="${arch}" go build -ldflags "-s -w"  -o "${output_dir}/kwcli" ../main.go
}

# Target OS/Architecture combinations
os_arch_list=("linux-amd64" "windows-amd64") # Add more as needed

# Build for each target
for target in "${os_arch_list[@]}"; do
    os_arch=(${target//-/ }) # Split into an array
    build_for_os_arch "${os_arch[0]}" "${os_arch[1]}" 
done

echo "Build complete!"
