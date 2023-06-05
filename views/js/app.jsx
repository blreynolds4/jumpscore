class App extends React.Component {
    render() {
        return (<Home />);
    }
}


class Home extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
          events: []
        };
        
        this.serverRequest = this.serverRequest.bind(this);
    }

    serverRequest() {
        $.get("http://localhost:3000/api/events", res => {
          this.setState({
            events: res
          });
        });
    }    

    render() {
      return (
        <div className="container">
            <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                <h1>Ski Jump Scoring</h1>
                <p>Basic scoring for club meets.</p>
                <div className="row">
                    <div className="container">
                        {this.state.events.map(function(event, i) {
                            return <Event key={i} event={event} />;
                        })}
                    </div>
                </div>
            </div>
        </div>
      )
    }
}

class Event extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        event: []
      };
      this.serverRequest = this.serverRequest.bind(this);
    }
      
    serverRequest(event) {
      $.post(
        "http://localhost:3000/api/events/" + event.id,
        res => {
          console.log("res... ", res);
          this.setState({ event: res });
          this.props.event = res;
        }
      );
    }
      
    render() {
      return (
        <div className="col-xs-4">
          <div className="panel panel-default">
            <div className="panel-heading">
              #{this.props.event.name}{" "}
            </div>
            <div className="panel-body">
                <a onClick={this.like} className="btn btn-default">
                </a>
            </div>
          </div>
        </div>
      )
    }
}

ReactDOM.render(<App />, document.getElementById('app'));