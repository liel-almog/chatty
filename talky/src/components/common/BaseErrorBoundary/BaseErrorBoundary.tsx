import { faTriangleExclamation } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Component, ReactNode, ComponentType } from "react";
import classes from "./base-error-boundary.module.scss";

interface Props {
  children?: ReactNode;
}

interface State {
  hasError: boolean;
  errorMessage: string;
}

export class BaseErrorBoundary extends Component<Props, State> {
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  //@ts-ignore
  constructor(props) {
    super(props);
    this.state = { errorMessage: "", hasError: false };
  }

  componentDidCatch(error: { name: string; message: string }) {
    this.setState({
      errorMessage: `שגיאה: ${error.message}`,
      hasError: true,
    });
  }

  // internet explorer doesn't support this project anyways.
  componentDidMount() {
    //checks what browser the user is using.
    const ie = false;
    if (ie) {
      this.setState({ errorMessage: "דפדפן לא נתמך: איטרנט אקספלורר" });
    }
    // internet explorer logic set error
  }

  render() {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    //@ts-ignore
    const { errorMessage } = this.state;
    if (errorMessage) {
      return (
        <div className={classes.container}>
          <FontAwesomeIcon icon={faTriangleExclamation} />
          <h1>אירעה שגיאה</h1>
          <p>{errorMessage}</p>
        </div>
      );
    } else {
      return <>{this.props.children}</>;
    }
  }
}
