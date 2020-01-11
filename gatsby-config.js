require('dotenv').config({
  path: `.env.${process.env.NODE_ENV}`,
});

module.exports = {
  siteMetadata: {
    title: `Worlds Best`,
    description: `Worlds best food, with passion and love`,
    author: `@masiuciszek`,
  },
  plugins: [
    `gatsby-plugin-react-helmet`,
    `gatsby-plugin-typescript`,
    `gatsby-plugin-styled-components`,
    `gatsby-plugin-transition-link`,
    `gatsby-plugin-playground`,
    {
      resolve: `gatsby-source-contentful`,
      options: {
        spaceId: 'sljqnyltboss',
        // Learn about environment variables: https://gatsby.dev/env-vars
        accessToken: 'FZUnzys7M1of78zinHJuFZgI-Q_LNTbc4GRe0xzGYz8',
      },
    },
    {
      resolve: `gatsby-source-filesystem`,
      options: {
        name: `images`,
        path: `${__dirname}/src/images`,
      },
    },
    `gatsby-transformer-sharp`,
    `gatsby-plugin-sharp`,
    {
      resolve: `gatsby-plugin-manifest`,
      options: {
        name: `worlds best`,
        shortName: `worlds`,
        startUrl: `/`,
        backgroundColor: `#fff`,
        themeColor: `#333`,
        display: `minimal-ui`,
        icon: `src/images/gatsby-icon.png`, // This path is relative to the root of the site.
      },
    },
  ],
};
