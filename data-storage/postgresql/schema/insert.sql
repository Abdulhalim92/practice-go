INSERT INTO public.teacher (create_time, firstname, lastname)
VALUES (NOW(),$1, $2)
    RETURNING id;