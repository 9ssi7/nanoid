## Getting Started

### 1. Clone the Repository

To contribute to this project, you need to clone the repository to your local machine.

```sh
git clone https://github.com/9ssi7/nanoid.git
cd nanoid
```

### 2. Run the Tests

Before making any changes, run the tests to ensure everything is working correctly.

```sh
go test ./...
```

### 3. Run the Benchmarks

To run the benchmarks and compare the performance of NanoID with Google UUID, use the following command:

```sh
go test -bench=. ./...
```

## Contributing

### 1. Create a Branch

Create a new branch for your feature or bugfix.

```sh
git checkout -b feature/my-new-feature
```

### 2. Make Your Changes

Implement your feature or bugfix and make sure to add tests.

### 3. Run Tests and Benchmarks

Before submitting your changes, run the tests and benchmarks to ensure everything is working correctly.

```sh
go test ./...
go test -bench=. ./...
```

### 4. Push Your Changes

Once you've committed your changes, push your branch to GitHub.

```sh
git push origin feature/my-new-feature
```

### 5. Create a Pull Request

Go to the [NanoID GitHub repository](https://github.com/9ssi7/nanoid) and create a new pull request from your branch.

## Support

If you encounter any issues or have questions about the project, please open an issue on GitHub.
