const std = @import("std");

fn caesarEncrypt(text: []const u8, shift: u8, allocator: std.mem.Allocator) ![]u8 {
    var result = try allocator.alloc(u8, text.len);

    for (text, 0..) |c, i| {
        if (c >= 'a' and c <= 'z') {
            result[i] = 'a' + (c - 'a' + shift) % 26;
        } else if (c >= 'A' and c <= 'Z') {
            result[i] = 'A' + (c - 'A' + shift) % 26;
        } else {
            result[i] = c; // caractères non alphabétiques inchangés
        }
    }

    return result;
}

fn caesarDecrypt(text: []const u8, shift: u8, allocator: std.mem.Allocator) ![]u8 {
    return caesarEncrypt(text, 26 - (shift % 26), allocator);
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;

    const message = "Bonjour Zig!";
    const shift: u8 = 3;

    const encrypted = try caesarEncrypt(message, shift, allocator);
    defer allocator.free(encrypted);

    const decrypted = try caesarDecrypt(encrypted, shift, allocator);
    defer allocator.free(decrypted);

    const stdout = std.io.getStdOut().writer();

    try stdout.print("Message original : {s}\n", .{message});
    try stdout.print("Message chiffré  : {s}\n", .{encrypted});
    try stdout.print("Message déchiffré: {s}\n", .{decrypted});
}

