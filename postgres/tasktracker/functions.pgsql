-- Create a task
CREATE OR REPLACE FUNCTION tasktracker.create_task(
    task_name   TEXT,
    description TEXT,
    owner_id    INTEGER
) RETURNS INT
    LANGUAGE plpgsql AS
$$
DECLARE
    _id INT;
BEGIN
    INSERT
        INTO
            tasktracker.task
            (name, description, owner_id)
        VALUES
            (task_name, description, owner_id)
        RETURNING id INTO _id;

    IF NOT found
        THEN
            RAISE EXCEPTION 'Could not create a task. Name:%; description:%; owner:%.',
                task_name, description, owner_id;
    END IF;

    RETURN _id;
END;
$$;

-- Get tasks by user
CREATE OR REPLACE FUNCTION tasktracker.get_tasks(
    user_id INT
) RETURNS TABLE (LIKE tasktracker.task)
    LANGUAGE plpgsql AS
$$
BEGIN
    RETURN QUERY
        SELECT *
            FROM
                tasktracker.task t
            WHERE
                t.owner_id = user_id;
END;
$$;

-- Update task
CREATE OR REPLACE PROCEDURE tasktracker.update_task(
    task_id         INT,
    new_name        TEXT,
    new_description TEXT,
    new_status      tasktracker.task_status,
    user_id         INT
)
    LANGUAGE plpgsql AS
$$
DECLARE
    update_query TEXT;
    updated      INT;
BEGIN
    update_query := 'WITH update_cte AS ('
                        || 'UPDATE tasktracker.task SET '
                        || concat_ws(
                            ', ',
                            CASE WHEN new_name IS NOT NULL THEN 'name = $1 ' END,
                            CASE WHEN new_description IS NOT NULL THEN 'description = $2 ' END,
                            CASE WHEN new_status IS NOT NULL THEN 'status = $3 ' END
                        )
                        || 'WHERE id = $4 AND owner_id = $5 '
                        || 'RETURNING 1'
        || ') SELECT count(*) FROM update_cte';

    EXECUTE update_query USING new_name, new_description, new_status, task_id, user_id INTO updated;

    IF updated = 0
        THEN
            RAISE EXCEPTION 'task not found';
    END IF;
END;
$$;

-- Delete a task
CREATE OR REPLACE PROCEDURE tasktracker.delete_task(
    task_id INT
)
    LANGUAGE plpgsql AS
$$
BEGIN
    DELETE
        FROM
            tasktracker.task
        WHERE
            id = task_id;
    IF NOT found
        THEN
            RAISE EXCEPTION 'task not found';
    END IF;
END;
$$;