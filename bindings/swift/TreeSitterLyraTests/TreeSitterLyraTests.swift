import XCTest
import SwiftTreeSitter
import TreeSitterLyra

final class TreeSitterLyraTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_lyra())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Lyra grammar")
    }
}
