# go-qmp

An experimental QMP client for Go.

## Example

```golang
import (
    "github.com/saracen/go-qmp"
    "github.com/saracen/go-qmp/api/qsd"
)

func createBlockDevice() error {
    client, err := qmp.Dial(ctx, "unix", "qmp.sock")
    if err != nil {
        return err
    }

    // register qemu storage daemon event factory
    qsd.RegisterEvents(client)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    cancel()

    // create listener and filter for certain event & job id
    events := client.Listen(ctx, func(e qmp.Event) bool {
        switch v := e.Event.(type) {
        case *qsd.JobStatusChangeEvent:
            return v.Id == "my-job-id"
        }

        return false
    })

    // create blockdev
    create := qsd.BlockdevCreate{
        JobId: "my-job-id",
        Options: qsd.BlockdevCreateOptions{
        Driver: qsd.BlockdevDriverFile,
            File: &qsd.BlockdevCreateOptionsFile{
                Filename:      "/my/raw.disk",
                Size:          10 * 1024 * 1024,
                Preallocation: qmp.ToPtr(qsd.PreallocModeFalloc),
            },
        },
    }

    if err := create.Execute(ctx, client); err != nil {
        return fmt.Errorf("executing create block device: %w", err)
    }

    // watch events
    var errs []error
    for event := range events {
        statusChange, ok := event.Event.(*qsd.JobStatusChangeEvent)
        if !ok {
            continue
        }

        switch statusChange.Status {
        default:
            continue

        case qsd.JobStatusAborting:
            jobInfo, err := qsd.QueryJobs{}.Execute(ctx, client)
            if err != nil {
                errs = append(errs, fmt.Errorf("querying jobs: %w", err))
            } else {
                // find the error for the job id
                for _, info := range jobInfo {
                    if info.Id == "my-job-id" {
                        errs = append(errs, errors.New(qmp.FromPtr(info.Error)))
                        break
                    }
                }
            }

            continue

        case qsd.JobStatusConcluded:
            err := qsd.JobDismiss{Id: "my-job-id"}.Execute(ctx, .client)
            if err != nil {
                errs = append(errs, fmt.Errorf("dismissing job: %w", err))
            }
        }

        break
    }

    if err := errors.Join(errs...); err != nil {
        return fmt.Errorf("creating block device: %w", err)
    }

    return nil
}
```
