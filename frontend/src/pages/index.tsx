import type { NextPage } from "next";
import { useQuery, useMutation } from "@apollo/client";

import { CreateTodoDocument, CreateTodoMutation, GetTodoDocument } from "../../graphql/dist/client";
import { GetTodoQuery } from "../../graphql/dist/client";
import { useState } from "react";
import { Button, Card, Container, Grid, Input, Spacer, Text } from "@nextui-org/react";
import { randomInt } from "crypto";


const Home: NextPage = () => {
  const { data, loading, refetch } = useQuery<GetTodoQuery>(GetTodoDocument);
  const [text, setText] = useState("");

  const [addTodo] = useMutation<CreateTodoMutation>(CreateTodoDocument);

  if (loading) return null;

  return (
    <Container>
      <Text h1 color="primary">ToDo App</Text>
      {data?.todos?.map((todo) => (
        <>
          <Card key={todo.id}>
            <Card.Body>
              <Text h3>{todo.content}</Text>
              <Text h6>{todo.id}</Text>
            </Card.Body>
          </Card>
          <Spacer y={0.5}/>
        </>
      ))}
      <Grid.Container gap={2} justify="center">
        <Grid xs={8}>
          <Input
            bordered
            placeholder="ToDoを入力してください"
            fullWidth
            type="text"
            id="message"
            name="message"
            onChange={(e) => { setText(e.target.value) }}
            value={text}
          />
        </Grid>
        <Grid xs={4}>
          <Button
            onClick={
              async () => {
                await addTodo({ variables: { content: text, userId: (Math.floor(Math.random() * 100)).toString() } });
                setText("");
                refetch()
              }
            }
          >
            ToDoを追加
          </Button>
        </Grid>
      </Grid.Container>


    </Container>
  );
};

export default Home;