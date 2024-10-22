db.createUser({
    user: 'user1',
    pwd: 'Qwerty123',
    roles: [
        {
            role: 'readWrite',
            db: 'users'
        }
    ]
});