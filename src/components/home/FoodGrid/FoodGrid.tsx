import * as React from 'react';
import { StyledFoodGrid, GridItem } from './Styles.foodgrid';
import { useStaticQuery, graphql } from 'gatsby';
import Img from 'gatsby-image';
import { SingleNodeImg } from '../../../types';

const GET_IMAGES = graphql`
  query {
    images: allFile(filter: { relativeDirectory: { eq: "homeGallery" } }) {
      edges {
        node {
          name
          sourceInstanceName
          relativePath
          childImageSharp {
            fluid(maxWidth: 400) {
              ...GatsbyImageSharpFluid_tracedSVG
            }
          }
        }
      }
    }
  }
`;

interface Props {
  // images: {
  //   edges: {
  //     node: {
  //       name: string;
  //       sourceInstanceName: string;
  //       relativePath: string;
  //       childImageSharp: {
  //         fluid?: FluidObject | FluidObject[] | undefined;
  //       };
  //     };
  //   }[];
  // };
  images: {
    edges: {
      node: SingleNodeImg;
    }[];
  };
}

const FoodGrid: React.FC<Props> = () => {
  const data: Props = useStaticQuery(GET_IMAGES);
  console.log('img', data);
  return (
    <StyledFoodGrid>
      {data.images.edges.map((x, idx) => {
        return (
          <GridItem key={idx}>
            <Img fluid={x.node.childImageSharp.fluid} />
            <p>{x.node.name}</p>
          </GridItem>
        );
      })}
    </StyledFoodGrid>
  );
};
export default FoodGrid;
