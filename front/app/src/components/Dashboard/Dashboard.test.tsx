import { render, screen } from "@testing-library/react"
import { rest } from "msw"
import { setupServer } from "msw/lib/node"
import { Provider } from "react-redux"
import { store } from "../../store"
import Dashboard from "./Dashboard"

describe("test for dashboard", () => {
  const handlers = [
    rest.get("http://localhost:3000/api/v1/lists", async (req, res, ctx) => {
      return res(
        ctx.status(200),
        ctx.json(
          [
            {
              "id": 1,
              "title": "l1",
              "user_id": 1,
              "created_at": "2022-09-25T01:39:45.152Z",
              "updated_at": "2022-09-25T01:39:45.152Z"
            },
            {
              "id": 2,
              "title": "l2",
              "user_id": 2,
              "created_at": "2022-09-25T01:40:00.215Z",
              "updated_at": "2022-09-25T01:40:00.215Z"
            },
            {
              "id": 3,
              "title": "l3",
              "user_id": 3,
              "created_at": "2022-09-25T01:40:21.426Z",
              "updated_at": "2022-09-25T01:40:21.426Z"
            },
            {
              "id": 4,
              "title": "l4",
              "user_id": 4,
              "created_at": "2022-09-25T01:40:43.730Z",
              "updated_at": "2022-09-25T01:40:43.730Z"
            }
          ]
        )
      )
    }),
    rest.get("http://localhost:3000/api/v1/cards", async (req, res, ctx) => {
      const list_id = req.url.searchParams.get("list_id")
      if (list_id !== "1") {
        return res(
          ctx.status(200)
        );
      }
      return res(
        ctx.status(200),
        ctx.json(
          [
            {
              "id": 1,
              "title": "t1",
              "body": "b1",
              "list_id": 1,
              "created_at": "2022-09-25T01:39:54.161Z",
              "updated_at": "2022-09-25T01:39:54.161Z"
            }
          ]
        )
      )
    })
  ]
  const server = setupServer(...handlers)

  beforeAll(() => server.listen())
  afterEach(() => server.resetHandlers())
  afterAll(() => server.close())
  test("renders list title", async () => {
    render(
      <Provider store={store}>
        <Dashboard />
      </Provider>
    );
    const l1 = await screen.findByText("l1");
    expect(l1).not.toBe(null);
  })
  test("renders card title", async () => {
    render(
      <Provider store={store}>
        <Dashboard />
      </Provider>
    );
    const t1 = await screen.findByText("t1");
    expect(t1).not.toBe(null);
  })

})