import React, {Component} from 'react';

export default class Welcome extends Component {
  render() {
      return (
          <div className="container">
              <div className="row">
                  <div className="col-md-8 col-md-offset-2">
                      <div className="panel panel-default">
                          <div className="panel-heading">Hello, World</div>

                          <div className="panel-body">
                              I am an example react component!
                          </div>
                      </div>
                  </div>
              </div>
          </div>
      );
  }
}