CREATE FUNCTION add_class_to_user(_user_id INT, _class_id INT)
   RETURNS SETOF classifications
   LANGUAGE PLPGSQL
AS $$
DECLARE
   rcd classifications;
BEGIN
   INSERT INTO user_classifications (user_id, class_id) VALUES (_user_id, _class_id);

   RETURN QUERY (
      SELECT c.* FROM classifications AS c INNER JOIN user_classifications AS uc ON c.class_id=uc.class_id
   );
END;$$;