module.exports = {
    setupFilesAfterEnv: ["./jest.setup.js"],
    moduleNameMapper: {
        '\\.(css|less)$': '<rootDir>/test/jest/__mocks__/styleMock.js',
        "^@components(.*)$": "<rootDir>/components$1",
        "^@pages(.*)$": "<rootDir>/pages$1",
        "^@hooks(.*)$": "<rootDir>/hooks$1",
    },
    testEnvironment: 'jsdom',
};