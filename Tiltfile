docker_compose('docker-compose.yml')

def ko_service(name, path):
    custom_build(
        name,
        # macOS / Linux / CI
        command='KO_DOCKER_REPO=$EXPECTED_IMAGE ko publish --local --bare --tags $EXPECTED_TAG ' + path,
        # Windows (cmd.exe) — avoids bash / WSL entirely
        command_bat='set KO_DOCKER_REPO=%EXPECTED_IMAGE%&& ko publish --local --bare --tags %EXPECTED_TAG% ' + path,
        deps=[path, 'go.mod', 'go.sum'],
    )

# Register services here:
ko_service('distributed-go/test-api', './cmd/TestAPI')
