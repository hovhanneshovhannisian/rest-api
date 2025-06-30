todo implementing messenger

1. Chat Rooms Table:
   Purpose: Manages different chat rooms or conversations within the application.
   Fields:
   room_id (INT, PRIMARY KEY, AUTO_INCREMENT): Unique identifier for each chat room.
   name (VARCHAR): Name of the chat room (e.g., "General", "Project Team").
   created_at (DATETIME, DEFAULT CURRENT_TIMESTAMP): Timestamp of when the chat room was created.
   type (ENUM): Indicates whether it's a one-on-one chat (e.g., "direct") or a group chat (e.g., "group").
2. Messages Table:
   Purpose: Stores the actual messages exchanged within the chat rooms.
   Fields:
   message_id (INT, PRIMARY KEY, AUTO_INCREMENT): Unique identifier for each message.
   room_id (INT, NOT NULL, FOREIGN KEY referencing ChatRooms.room_id): Foreign key linking the message to its chat room.
   sender_id (INT, NOT NULL, FOREIGN KEY referencing Users.user_id): Foreign key linking the message to the sender.
   content (TEXT): The actual text content of the message.
   created_at (DATETIME, DEFAULT CURRENT_TIMESTAMP): Timestamp of when the message was sent.
   is_read (BOOLEAN, DEFAULT FALSE): Indicates whether the message has been read by the recipient(s).
3. Chat memeber table:
   to store the members of the chats
   fields:
   id:
   user_id: foreign key users.userid
   room_id: foreign key rooms.room_id
   role: enum either member or admin
   join_at: time
