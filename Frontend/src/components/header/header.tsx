import dumbbellIco from "@/../public/svg/dumbbell.svg";
import style from "./header.module.css";
import Image from "next/image";
import Link from "next/link";

export default function Header() {
  return (
    <header className={style.header}>
      <div className={style.inner}>
        <div className={style.brand}>
          <Link className={style.brandLink} href="/">
            <Image
              className={style.logo}
              src={dumbbellIco}
              alt="Иконка в виде гантели"
            />
            <b className={style.title}>Calisthetics</b>
          </Link>
        </div>

        <nav className={style.nav}>
          <NavItem content="Упражнения" href="/exercises" />
          <NavItem content="Программы" href="/programs" />
          {true ? (
            <NavItem content="Профиль" href="/profile" />
          ) : (
            <NavItem content="Авторизация" href="/login" />
          )}
        </nav>
      </div>
    </header>
  );
}

function NavItem({ content, href }: { content: string; href: string }) {
  return (
    <Link href={href} className={style.navItem}>
      {content}
    </Link>
  );
}
