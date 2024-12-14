import styles from "./footer.module.css";
import Link from "next/link";

const footerContent = [
  {
    key: 0,
    title: "Навигация",
    links: [
      { name: "Главная", link: "/", target: "_self" },
      { name: "Упражнения", link: "/exercises", target: "_self" },
      { name: "Программы", link: "/programs", target: "_self" },
      { name: "Профиль", link: "/profile", target: "_self" },
    ],
  },
  {
    key: 1,
    title: "Контакты",
    links: [
      {
        name: "Github",
        link: "https://github.com/DeSheans",
        target: "_blank",
      },
    ],
  },
];

export default function Footer() {
  return (
    <div className={styles.footer}>
      <div className={styles.inner}>
        {footerContent.map((col) => (
          <Column key={col.key} title={col.title} links={col.links} />
        ))}
      </div>
    </div>
  );
}

function Column({
  title,
  links,
}: {
  title: string;
  links: { name: string; link: string; target: string }[];
}) {
  return (
    <div className={styles.column}>
      <div className={styles.columnTitle}>{title}</div>
      {links.map((link) => (
        <Link
          key={link.name}
          className={styles.columnItem}
          href={link.link}
          target={link.target}
        >
          {link.name}
        </Link>
      ))}
    </div>
  );
}
