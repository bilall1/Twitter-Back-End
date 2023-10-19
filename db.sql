CREATE TABLE users (
	Id serial PRIMARY KEY,
	Email text NOT NULL UNIQUE,
	Password text,
	third_party boolean DEFAULT FALSE,
	D_o_b date,
	first_name text,
	last_name text,
	Profile text
);

CREATE TABLE tweets (
	Id serial PRIMARY KEY,
	Content text,
	User_id int REFERENCES users(Id) ON DELETE CASCADE,
	Link text
);

CREATE TABLE User_Followers (
	Id serial PRIMARY KEY,
	User_Id int REFERENCES users(Id) ON DELETE CASCADE,
	Follower_Id int REFERENCES users(Id) ON DELETE CASCADE
);

CREATE TABLE Tweets_Likes (
	Id serial PRIMARY KEY,
	Tweet_Id int REFERENCES tweets(Id) ON DELETE CASCADE,
	User_Id int REFERENCES users(Id) ON DELETE CASCADE
);

CREATE TABLE Message (
	Id serial PRIMARY KEY,
	Sender_Id int REFERENCES user(Id) ON DELETE CASCADE,
	Reciever_Id int REFERENCES user(Id) ON DELETE CASCADE,
	Message_Type text,
	Created_At TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	Status text,
	Content text
);

CREATE TABLE Conversations (
	Id serial PRIMARY KEY,
	Participant1 int REFERENCES users(Id) ON DELETE CASCADE,
	Participant2 int REFERENCES users(Id) ON DELETE CASCADE,
	Last_Chat TIMESTAMP WITHOUT TIME ZONE ,
	Last_Message text
);



CREATE TABLE User_Status (
	Id serial PRIMARY KEY,
	User_Id int REFERENCES users(Id) ON DELETE CASCADE,
	Last_Active TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	Status text
);


CREATE OR REPLACE FUNCTION update_or_insert_conversation()
RETURNS TRIGGER AS $$
BEGIN
    -- Check if a conversation between SenderId and ReceiverId already exists
    IF EXISTS(SELECT 1 FROM Conversations WHERE (Participant1 = NEW.Sender_Id AND Participant2 = NEW.Reciever_Id) OR (Participant1 = NEW.Reciever_Id AND Participant2 = NEW.Sender_Id)) THEN
        -- Update the Last_Chat and Last_Message columns of the existing conversation
        UPDATE Conversations 
        SET Last_Chat = NEW.Created_At, 
            Last_Message = NEW.Content
        WHERE (Participant1 = NEW.Sender_Id AND Participant2 = NEW.Reciever_Id) OR (Participant1 = NEW.Reciever_Id AND Participant2 = NEW.Sender_Id);
    ELSE
        -- If no conversation exists, create a new one
        INSERT INTO Conversations (Participant1, Participant2, Last_Chat, Last_Message)
        VALUES (NEW.Sender_Id, NEW.Reciever_Id, NEW.Created_At, NEW.Content);
    END IF;

    RETURN NEW; -- Return the new row for the INSERT to continue
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER tr_after_insert_message
AFTER INSERT ON Messages
FOR EACH ROW
EXECUTE FUNCTION update_or_insert_conversation();
