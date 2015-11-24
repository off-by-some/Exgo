### TODO:
[ ] Session Creation (Create session token, map t user id, set cookie)

[ ] Session Tracking (Lookup user id with session token so resources can use it)

[ ] End Session (Revoke the session token, should automatically happen after a couple hours if the user doesn't explicitly log out)

[ ] Error Handling

Note: Gorilla apparently has a nice way of managing cookies and sessions, should probably use that.
