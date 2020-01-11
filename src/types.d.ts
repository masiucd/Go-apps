import { FluidObject } from 'gatsby-image';

export type SingleNodeImg = {
  name: string;
  sourceInstanceName: string;
  relativePath: string;
  childImageSharp: {
    fluid?: FluidObject | FluidObject[] | undefined;
  };
};
