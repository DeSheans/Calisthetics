import Header from "@/components/header/header";
import Footer from "@/components/footer/footer";
import styles from "./layout.module.css";
import "@/styles/global.css";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ru">
      <body>
        <div className={styles.container}>
          <header className={styles.header}>
            <Header />
          </header>
          <main className={styles.main}>{children}</main>
          <footer className={styles.footer}>
            <Footer />
          </footer>
        </div>
      </body>
    </html>
  );
}
