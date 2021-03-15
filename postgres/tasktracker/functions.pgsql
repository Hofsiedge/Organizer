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
            RAISE EXCEPTION 'Task % not found', task_id;
    END IF;
END;
$$;

-- TODO
