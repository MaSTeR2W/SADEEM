CREATE VIEW joined_users_classifications AS
    SELECT 
        u.*,
        ARRAY_AGG(
            JSON_BUILD_OBJECT(
                'classId',
                c.class_id,
                'name',
                c.name,
                "enabled",
                c.enabled
            )
        ) AS classifications
    FROM users AS u
    LEFT JOIN user_classifications AS uc ON u.user_id=uc.user_id
    LEFT JOIN classifications AS c ON c.class_id=uc.class_id
    GROUP BY u.user_id;