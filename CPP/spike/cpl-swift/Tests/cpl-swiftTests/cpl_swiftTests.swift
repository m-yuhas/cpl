import XCTest
@testable import cpl_swift

class cpl_swiftTests: XCTestCase {
    func testExample() {
        // This is an example of a functional test case.
        // Use XCTAssert and related functions to verify your tests produce the correct results.
        XCTAssertEqual(cpl_swift().text, "Hello, World!")
    }


    static var allTests : [(String, (cpl_swiftTests) -> () throws -> Void)] {
        return [
            ("testExample", testExample),
        ]
    }
}
