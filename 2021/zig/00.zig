const std = @import("std");
const io = std.io;
const fmt = std.fmt;

pub fn main() !void {
    const stdin = io.getStdIn().reader();
    const allocator = std.heap.page_allocator;
    var depthHistory = std.ArrayList(u8).init(allocator);
    defer depthHistory.deinit();

    var buf: [1024]u8 = undefined;
    while (try stdin.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        const depth = fmt.parseUnsigned(u8, line, 10) catch {
            continue;
        };
        std.debug.print("{d}\n", .{depth});
    }


    //while (true){
    //    var line_buf: [20]u8 = undefined;
    //    const amt = try stdin.read(&line_buf);
    //    if (amt == line_buf.len) {
    //        std.debug.print("Input too long: {d}!\n", .{amt});
    //        continue;
    //    }
    //    const line = std.mem.trimRight(u8, line_buf[0..amt], "\r\n");
    //        continue;
    //    };
    //    try depthHistory.append(depth);
    //}
}
