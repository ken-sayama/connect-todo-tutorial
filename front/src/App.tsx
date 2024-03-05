import './App.css'
import {Client} from "./Client.tsx";
import {GetTodo} from "./Todo.tsx";

function App() {
  const baseUrl = "http://localhost:8080";

  return (
    <Client baseUrl={baseUrl}>
        <h1>Hello</h1>
        <GetTodo id="hoge" />
    </Client>
  )
}

export default App
