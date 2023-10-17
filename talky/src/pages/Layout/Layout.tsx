import { Outlet, ScrollRestoration } from "react-router-dom";
import { Footer } from "../../components/Footer";
import { Header } from "../../components/Header";
import classes from "./layout.module.scss";

export const Layout: React.FC = () => {
  return (
    <>
      <ScrollRestoration
        getKey={(location, matches) => {
          let match = matches.find((m) => (m.handle as any)?.scrollMode);
          if ((match?.handle as any)?.scrollMode === "pathname") {
            return location.pathname;
          }

          return location.key;
        }}
      />
      <div className={classes.container}>
        <Header />
        <main>
          <Outlet />
        </main>
        <Footer />
      </div>
    </>
  );
};
