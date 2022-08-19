const std = @import("std");
const io = std.io;
const fmt = std.fmt;

pub fn main() !void {
    const allocator = std.heap.page_allocator;

    // Read from stdin, line by line, into ArrayList of u16s
    var buf: [1024]u8 = undefined;
    const stdin = io.getStdIn().reader();
    var depthReport = std.ArrayList(u16).init(allocator);
    defer depthReport.deinit();
    while (try stdin.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        const depth = fmt.parseUnsigned(u16, line, 10) catch {
            std.debug.print("PARSE ERR: {s}\n", .{line});
            continue;
        };
        try depthReport.append(depth);
    }

    // Create derivative 3 sweep array
    var depth3 = std.ArrayList(u16).init(allocator);
    defer depth3.deinit();
    const dR = depthReport.items;
    for (dR) |_,i| {
        if (i > 1) {
            const sweep = dR[i-2] + dR[i-1] + dR[i];
            try depth3.append(sweep);
        }
    }

    // Count number of increases in depth3 list
    var inc: u16 = 0;
    const d3 = depth3.items;
    for (d3) |d,i| {
        if (i > 0) {
            const prev = d3[i-1];
            if (prev < d){
                inc += 1;
            }
        }
    }
    std.debug.print("{d}\n", .{inc});
}
