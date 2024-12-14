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
          <div className={styles.header}>
            <Header />
          </div>
          <main className={styles.main}>{children}</main>
          <div className={styles.footer}>
            <Footer />
          </div>
        </div>
      </body>
    </html>
  );
}
