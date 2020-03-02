import React, { Component } from "react";
import fire from "./fire";
import Button from "@material-ui/core/Button";

class Home extends Component {
  constructor(props) {
    super(props);
    this.logout = this.logout.bind(this);
  }

  logout() {
    fire.auth().signOut();
  }

  render() {
    return (
      <div>
        <h1>Hello</h1>
        <Button onClick={this.logout} color="primary">
          logout
        </Button>
      </div>
    );
  }
}

export default Home;
