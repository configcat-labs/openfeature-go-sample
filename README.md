# Courses App

#### (Optional) In case the blog post is already published, please add: [Read the blog post here](https://configcat.com/blog/)

This is a simple app that demonstrates how to use [ConfigCat](https://configcat.com) with [OpenFeature](https://openfeature.dev/) in Go. The app has a REST API with a single endpoint that returns a predefined list of courses. We can enable or disable the endpoint with a feature flag in ConfigCat.

**NOTE:** This branch contains the complete code with OpenFeature integrated into the app. Use `git checkout starter` to switch to the `starter` branch to see the state of the app before OpenFeature was installed.

## Build & Run

## Prerequisites
- A ConfigCat account.
- Go 1.21.4+.
- Git
- A Terminal.
- A tool to make HTTP requests. Examples: curl, Postman, Thunder Client, e.t.c.
- Intermediate knowledge of Go and basic knowledge of Git.

## Instructions for Running
Follow these steps to run the project on your computer:
- Open your terminal and clone this repository with this command:

`git clone https://github.com/configcat-labs/openfeature-go-sample.git`

- Navigate into the repository's folder:

`cd openfeature-go-sample`

- Open `main.go` and replace "YOUR-SDK-KEY" with your ConfigCat SDK key on this line:

`ccProvider := provider.NewProvider(configcat.NewClient("YOUR-SDK-KEY"))`

- Run the app:

`go run main.go`

- Visit `http://localhost:8000/courses`
- Toggle the feature flag in your ConfigCat dashboard to enable and disable the endpoint.

## Learn more

Useful links to technical resources.
- [ConfigCat OpenFeature Go Provider](https://github.com/open-feature/go-sdk-contrib/tree/main/providers/configcat) - check out the Go provider's repository
- [ConfigCat Go SDK Documentation](https://configcat.com/docs/sdk-reference/go) - learn more about ConfigCat's Go SDK
- [ConfigCat Java provider for OpenFeature](https://github.com/open-feature/java-sdk-contrib/tree/main/providers/configcat) - check out ConfigCat's Java Provider
- [ConfigCat JavaScript provider for OpenFeature](https://github.com/open-feature/js-sdk-contrib/tree/main/libs/providers/config-cat) - check out ConfigCat's JavaScript Provider
- [OpenFeature website](https://openfeature.dev/) - learn more about OpenFeature


[**ConfigCat**](https://configcat.com) also supports many other frameworks and languages. Check out the full list of supported SDKs [here](https://configcat.com/docs/sdk-reference/overview/).

You can also explore other code samples for various languages, frameworks, and topics here in the [ConfigCat labs](https://github.com/configcat-labs) on GitHub.

Keep up with ConfigCat on [Twitter](https://twitter.com/configcat), [Facebook](https://www.facebook.com/configcat), [LinkedIn](https://www.linkedin.com/company/configcat/), and [GitHub](https://github.com/configcat).

## Author
[Zayyad Muhammad Sani](https://github.com/Z-MS)

## Contributions
Contributions are welcome!