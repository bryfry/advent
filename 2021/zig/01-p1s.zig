const std = @import("std");
const io = std.io;
const fmt = std.fmt;

pub fn main() !void {
    const stdin = io.getStdIn().reader();
    const allocator = std.heap.page_allocator;
    var depthHistory = std.ArrayList(u16).init(allocator);
    defer depthHistory.deinit();

    var buf: [1024]u8 = undefined;
    while (try stdin.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        const depth = fmt.parseUnsigned(u16, line, 10) catch {
            std.debug.print("PARSE ERR: {s}\n", .{line});
            continue;
        };
        try depthHistory.append(depth);
    }
    var inc: u16 = 0;
    for (depthHistory.items) |d,i| {
        if (i > 0) {
            const prev = depthHistory.items[i-1];
            if (prev < d){
                inc += 1;
            }
        }
    }
    std.debug.print("{d}\n", .{inc});
}
