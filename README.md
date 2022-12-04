# byFood

Every question seperated by folders. Of course, `client-app` is on another folder.

## Usage

Firstly, in this workspace, we have more than one module. If you have errors when opening the workspace in `VSCode`, you can change the Go extension settings, in order to allow gopls to look for multiple modules in the workspace. Just add the following to your `settings.json`:

```json
"gopls": {
    "experimentalWorkspaceModule": true,
}
```

Now it should be fine.

Every question accept `Q4` have tests for their own.

To run those tests, you need to go to their folder.

For example:

```bash
cd Q1
go test -v ./...
```

This runs tests for `Q1`.

Now, for `Q4`, `PostgreSQL` has been used. Required database to make this work is in `Q4/pkg/persistence/db.go` file. It is called `byFood` in there. You need to create it manually. You can change it if you want of course. There are also connection informations in there, you need to change those for app to work.

App uses the default `public` schema. When yo start the Go, the `user` table will be created automatically.

You are all set!

Lastly, you need to run Go and React.

Go to `Q4/cmd/main/main.go` and run

`go run .`

For the client side, go to `client-app`

```
npm install
npm start
```

You are good to go!
