import type { NextPage } from "next";
import { useQuery, useMutation } from "@apollo/client";

import { CreateTodoDocument, CreateTodoMutation, GetTodoDocument } from "../../graphql/dist/client";
import { GetTodoQuery } from "../../graphql/dist/client";
import { useState } from "react";


const Home: NextPage = () => {
  const { data, loading, refetch } = useQuery<GetTodoQuery>(GetTodoDocument);
  const [text, setText] = useState("");

  const [addTodo] = useMutation<CreateTodoMutation>(CreateTodoDocument);

  if (loading) return null;

  return (
    <div style={{ margin: "0 auto", width: "1000px" }}>
      {data?.todos?.map((todo) => (
        <div key={todo.id}>
          <h1>{todo.text}</h1>
          <p>id:{todo.id}</p>
          <p>text:{todo.text}</p>
        </div>
      ))}
      <input
        type="text"
        id="message"
        name="message"
        onChange={(e) => { setText(e.target.value) }}
        value={text}
      />
      <button onClick={async () => { await addTodo({ variables: { text: text } }); setText(""); refetch() }}>Add Todo</button>
    </div>
  );
};

export default Home;