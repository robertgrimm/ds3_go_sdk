package commands

import (
    "io"
    "os"
    "path"
    "errors"
    "github.com/SpectraLogic/ds3_go_sdk/ds3"
    "github.com/SpectraLogic/ds3_go_sdk/ds3/models"
)

func bulkGet(client *ds3.Client, args *Arguments) error {
    // Validate arguments.
    if args.Bucket == "" {
        return errors.New("Must specify a bucket name when doing a bulk_get.")
    }

    // Determine the objects that we need to queue a bulk get for.
    objects, err := getBucketObjects(client, args)
    if err != nil {
        return err
    }

    var ds3GetObjects []models.Ds3GetObject
    for _, obj := range objects {
        ds3GetObjects = append(ds3GetObjects, models.NewDs3GetObject(obj.Name))
    }

    // Run request.
    response, err := client.GetBulkJobSpectraS3(models.NewGetBulkJobSpectraS3RequestWithPartialObjects(args.Bucket, ds3GetObjects))
    if err != nil {
        return err
    }

    // Handle the responses in parallel
    return handleBulkResponse(response.MasterObjectList.Objects, buildBulkHandler(client, args.Bucket))
}

// Returns a function that gets an object from a bulk get.
func buildBulkHandler(client *ds3.Client, bucketName string) bulkHandler {
    return func(obj models.BulkObject) error {
        // Perform the request.
        response, requestErr := client.GetObject(models.NewGetObjectRequest(bucketName, *obj.Name))
        if requestErr != nil {
            return requestErr
        }
        defer response.Content.Close()

        // Get a file to write to.
        file, fileErr := ensureDirectoryAndOpenFile(*obj.Name)
        if fileErr != nil {
            return fileErr
        }
        defer file.Close()

        // Copy the request stream to the file.
        _, copyErr := io.Copy(file, response.Content)
        return copyErr
    }
}

func ensureDirectoryAndOpenFile(destination string) (*os.File, error) {
    // Get the current directory so we can use the same permissions when
    // creating directories inside of it.
    currentDir, dirErr := os.Stat(".")
    if dirErr != nil {
        return nil, dirErr
    }

    // Create the directory where we're going to store the file.
    if mkdirErr := os.MkdirAll(path.Dir(destination), currentDir.Mode()); mkdirErr != nil {
        return nil, mkdirErr
    }

    // Open the file to write.
    return os.Create(destination)
}

