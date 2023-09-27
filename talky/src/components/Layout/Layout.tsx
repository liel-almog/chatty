import { Outlet } from 'react-router-dom';
import { Footer } from '../Footer';
import { Header } from '../Header';
import classes from './layout.module.scss';

export const Layout: React.FC = () => {


    return (
        <div className={classes.container}>
            <Header />
            <main>
                <Outlet />
            </main>
            <Footer />
        </div>
    )
}