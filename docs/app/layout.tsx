import Script from "next/script";
import { getPageMap } from "nextra/page-map";
import { Footer, Layout, Navbar } from "nextra-theme-docs";
import "nextra-theme-docs/style.css";

import "../globals.css";
import Analytics from "./analytics";

export const metadata = {
	title: {
		template: "%s – Ellie CLI",
		default: "Ellie CLI",
	},
	metadataBase: new URL("https://goprisma.org"),
	description:
		"Ellie CLI is an auto-generated and fully type-safe database client",
	applicationName: "Ellie CLI",
	icons: [
		{
			rel: "icon",
			url: "/icon.png",
			type: "image/png",
		},
	],
};

export default async function RootLayout({ children }) {
	return (
		<html lang="en" dir="ltr" suppressHydrationWarning>
			<head>
				<Script
					src="https://www.googletagmanager.com/gtag/js?id=G-SV698BGGW8"
					strategy="afterInteractive"
				/>
				<Script id="google-analytics" strategy="afterInteractive">
					{`
						window.dataLayer = window.dataLayer || [];
						function gtag(){dataLayer.push(arguments);}
						gtag('js', new Date());

						gtag('config', 'G-SV698BGGW8');
					`}
				</Script>
			</head>
			<body>
				<Analytics>
					<Layout
						footer={
							<Footer>
								<span>
									All source code and content licensed under&nbsp;
									<a
										href="https://github.com/gooferOrm/goofer/blob/main/LICENSE"
										target="_blank"
									>
										Apache 2.0
									</a>
								</span>
							</Footer>
						}
						navbar={
							<Navbar
								logo={<div>Ellie CLI</div>}
								chatLink="https://discord.gg/er3ZbmYHDk"
								projectLink="https://github.com/beyondEllie/ellie"
							/>
						}
						editLink="Edit this page on GitHub"
						docsRepositoryBase="https://github.com/beyondEllie/ellie/tree/main/docs"
						sidebar={{ defaultMenuCollapseLevel: 1 }}
						pageMap={await getPageMap()}
					>
						{children}
					</Layout>
				</Analytics>
			</body>
		</html>
	);
}
