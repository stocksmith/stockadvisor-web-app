import React, { Component } from "react";
import fire from "./fire";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";

class Login extends Component {
  constructor(props) {
    super(props);
    this.login = this.login.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.signup = this.signup.bind(this);
    this.state = {
      email: "",
      password: ""
    };
  }

  handleChange(e) {
    //console.log(e.target.name);
    this.setState({ [e.target.name]: e.target.value });
  }

  login(e) {
    e.preventDefault();
    fire
      .auth()
      .signInWithEmailAndPassword(this.state.email, this.state.password)
      .then(u => {})
      .catch(error => {
        console.log(error);
      });
  }

  signup(e) {
    console.log("meow");
    e.preventDefault();
    fire
      .auth()
      .createUserWithEmailAndPassword(this.state.email, this.state.password)
      .then(u => {})
      .then(u => {
        console.log(u);
      })
      .catch(error => {
        console.log(error);
      });
  }
  render() {
    return (
      <div className="col-md-6">
        <form>
          <label for="exampleInputEmail1">Email address</label>
          <TextField
            value={this.state.email}
            onChange={this.handleChange}
            type="email"
            name="email"
            class="form-control"
            id="exampleInputEmail1"
            aria-describedby="emailHelp"
            placeholder="Enter email"
          />

          <div class="form-group">
            <label for="exampleInputPassword1">Password</label>
            <TextField
              value={this.state.password}
              onChange={this.handleChange}
              type="password"
              name="password"
              class="form-control"
              id="exampleInputPassword1"
              placeholder="Password"
            />
          </div>
          <Button
            type="submit"
            onClick={this.login}
            variant="contained"
            color="primary"
          >
            Login
          </Button>
          <Button
            onClick={this.signup}
            style={{ marginLeft: "25px" }}
            variant="contained"
            color="primary"
          >
            Signup
          </Button>
        </form>
      </div>
    );
  }
}
export default Login;
