package cmd

import (
  "fmt"
  "os"

  kld "github.com/kaleido-io/kaleido-sdk-go/kaleido"
  "github.com/spf13/viper"
  "gopkg.in/resty.v1"
)

func getNewClient() kld.KaleidoClient {
  return kld.NewClient(viper.GetString("api.url"), viper.GetString("api.key"))
}

func validateName() {
  if name == "" {
    fmt.Println("Missing required parameter: name")
    os.Exit(1)
  }

  return
}

func validateProvider() {
  if provider == "" {
    fmt.Println("Missing required parameter: provider (quorum or geth)")
    os.Exit(1)
  }

  return
}

func validateConsensus() {
  if consensus == "" {
    fmt.Println("Missing required parameter: consensus (raft or ibft for quorum, or poa for geth)")
    os.Exit(1)
  }

  return
}

func validateConsortiumId(resourceName string, isCreate ...bool) {
  if consortiumId == "" {
    if len(isCreate) == 0 || isCreate[0] {
      fmt.Printf("Missing required parameter: --consortiumId for the consortium to add the new %s to\n", resourceName)
    } else {
      fmt.Printf("Missing required parameter: --consortiumId for the consortium to delete the %s from\n", resourceName)
    }

    os.Exit(1)
  }
}

func validateApiKeyId(resourceName string, isGet ...bool) {
  if apiKeyId == "" {
    if len(isGet) == 0 || isGet[0] {
      fmt.Printf("Missing required parameter: --apiKeyId to get the api key %s to\n", resourceName)
    } else {
      fmt.Printf("Missing required parameter: --apiKeyId to delete the api key the %s from\n", resourceName)
    }

    os.Exit(1)
  }
}

func validateEnvironmentId(resourceName string, isCreate ...bool) {
  if environmentId == "" {
    if len(isCreate) == 0 || isCreate[0] {
      fmt.Printf("Missing required parameter: --environmentId for the environment to add the new %s to\n", resourceName)
    } else {
      fmt.Printf("Missing required parameter: --environmentId for the environment to delete the %s from\n", resourceName)
    }

    os.Exit(1)
  }
}

func validateMembershipId(resourceName string, isCreate ...bool) {
  if membershipId == "" {
    if len(isCreate) == 0 || isCreate[0] {
      fmt.Printf("Missing required parameter: --membershipId for the membership to add the new %s to\n", resourceName)
    } else {
      fmt.Printf("Missing required parameter: --membershipId for the membership to delete the %s from\n", resourceName)
    }

    os.Exit(1)
  }
}

func validateDeleteId(resourceName string) {
  if deleteId == "" {
    fmt.Printf("Missing required parameter: --id for the %s to delete\n", resourceName)
    os.Exit(1)
  }

  return
}

func validateGetResponse(res *resty.Response, err error, resourceName string) {
  if res.StatusCode() != 200 {
    fmt.Printf("Could not retrieve %s. Status code: %d.", resourceName, res.StatusCode())
  }

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("\n%+v\n", res)
  }
}

func validateCreationResponse(res *resty.Response, err error, resourceName string) {
  if res.StatusCode() != 201 {
    fmt.Printf("Could not create %s. Status code: %d.", resourceName, res.StatusCode())
  }

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("\n%+v\n", res)
  }
}

func validateDeletionResponse(res *resty.Response, err error, resourceName string) {
  if res.StatusCode() != 202 {
    fmt.Printf("%s deletion failed. Status code: %d\n", resourceName, res.StatusCode())
  }

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("\n%+v\n", res)
  }
}