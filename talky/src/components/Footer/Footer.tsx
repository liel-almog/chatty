import packageJson from "../../../package.json";
import { ENVIRONMENT_SETTINGS } from "./environment.settings";
import classes from "./footer.module.scss";
export interface FooterProps {}

export const Footer = () => {
  return (
    <footer className={classes.container}>
      <EnvironmentIndicator />
      <span>פותח ע"י ליאל אלמוג</span>
      <VersionIndicator />
    </footer>
  );
};

const VersionIndicator = () => <span role="version">{packageJson.version}</span>;

const EnvironmentIndicator = () => {
  const labels = new Map(Object.entries(ENVIRONMENT_SETTINGS));

  const BlinkingDot = () => <span className={classes.dot} />;

  const Body = ({ label }: { label: string }) => {
    return (
      <span role="environment">
        <BlinkingDot />
        {"סביבת " + label}
      </span>
    );
  };

  if (labels.has(window.location.pathname)) {
    return <Body label={labels.get(window.location.pathname) as string} />;
  }

  return <Body label={"פיתוח מקומית"} />;
};
