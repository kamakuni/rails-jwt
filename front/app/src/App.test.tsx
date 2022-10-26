import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';
import { Provider } from 'react-redux';
import { store } from './store';
import { act } from 'react-dom/test-utils';
import { setupServer } from "msw/node";
import { rest } from "msw";

describe("test for app", () => {

  const handlers = [
    rest.get("http://localhost:3000/api/v1/login", (req, res, ctx) => {
      const { session_id } = req.cookies;
      const isValid = (session_id: string) => session_id === "valid"
      if (isValid(session_id)) {
        return res(
          ctx.status(200),
          ctx.json({
            message: "Users is authorized."
          })
        )
      } else {
        return res(
          ctx.status(401),
          ctx.json({
            message: "Users is unauthorized."
          })
        )
      }
    }),
    rest.post("http://localhost:3000/api/v1/login", async (req, res, ctx) => {
      const isValid = (email: string, password: string) => email === "valid" && password === "valid"
      const data = await req.json()
      if (isValid(data.email, data.password)) {
        return res(
          ctx.status(200),
          ctx.cookie("_session_id", "valid", {
            httpOnly: true,
            path: '/'
          }),
          ctx.json({
            message: "Users is authorized."
          })
        )
      } else {
        return res(
          ctx.status(400),
          ctx.json({
            message: "Users is unauthorized."
          })
        )
      }

    })
  ]

  const server = setupServer(...handlers)

  beforeAll(() => server.listen())
  afterEach(() => server.resetHandlers())
  afterAll(() => server.close())

  test('renders home link', () => {
    render(
      <Provider store={store}>
        <App />
      </Provider>
    );
    const anchorElement = screen.getAllByText("Home").filter((el) => { return el.nodeName === "A" })[0]
    expect(anchorElement).toHaveAttribute("href", "/")
  });

  test('renders dashboard link', () => {
    render(
      <Provider store={store}>
        <App />
      </Provider>
    );
    const anchorElement = screen.getAllByText("Dashboard").filter((el) => { return el.nodeName === "A" })[0]
    expect(anchorElement).toHaveAttribute("href", "/dashboard")
  });

  test('renders login link', () => {
    render(
      <Provider store={store}>
        <App />
      </Provider>
    );
    const anchorElement = screen.getAllByText("Login").filter((el) => { return el.nodeName === "A" })[0]
    expect(anchorElement).toHaveAttribute("href", "/login")
  });

  test('renders signup link', () => {
    render(
      <Provider store={store}>
        <App />
      </Provider>
    );
    const anchorElement = screen.getAllByText("Signup").filter((el) => { return el.nodeName === "A" })[0]
    expect(anchorElement).toHaveAttribute("href", "/signup")
  });

  test('renders home contents', () => {
    render(
      <Provider store={store}>
        <App />
      </Provider>
    );
    const anchorElement = screen.getAllByText("Home").filter((el) => { return el.nodeName === "A" })[0]
    act(() => {
      anchorElement.click();
    })
    const h2Element = screen.getAllByText("Home").filter((el) => { return el.nodeName === "H2" })[0]
    expect(h2Element).toHaveTextContent("Home")
  });

  test('renders unauthorized message after clicking /dashboard', async () => {
    render(
      <Provider store={store}>
        <App />
      </Provider>
    );
    const anchorElement = screen.getAllByText("Dashboard").filter((el) => { return el.nodeName === "A" })[0]
    await act(async () => {
      anchorElement.click();
    })
    const divElement = await screen.findByText("Users is unauthorized.");
    expect(divElement).not.toBe(null)
  });

  test('renders login form after clicking /login', async () => {
    render(
      <Provider store={store}>
        <App />
      </Provider>
    );
    const anchorElement = screen.getAllByText("Login").filter((el) => { return el.nodeName === "A" })[0]
    await act(async () => {
      anchorElement.click();
    })
    const labelElement = await screen.findByText("Email")
    expect(labelElement).not.toBe(null)
  })

  test('renders signup form after clicking /signup', async () => {
    render(
      <Provider store={store}>
        <App />
      </Provider>
    );
    const anchorElement = screen.getAllByText("Signup").filter((el) => { return el.nodeName === "A" })[0]
    await act(async () => {
      anchorElement.click();
    })
    const labelElement = await screen.findByText("Email")
    expect(labelElement).not.toBe(null)
  })

})