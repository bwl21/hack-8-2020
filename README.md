# hack-8-2020

This repository is for a session with bwl21 and 32leaves to learn

* Graphql
* React
* Go


# the usecases

Implement zupfnotenmanager an application to compile projects out of a given set of abc files.

1. One might have an arbitrary number of assets represented as abc files distributed among several "source". folders. every ABC file has metadata such as
   * Filename
   * Title
   * Genre
   * available extracts
   * preseleted extracts
   * rate of difficulty
   * rate of playability in an ensemble
   * copyright music
   * copyright words
   * reference copy indicator
   
2. out of this asses one might want to compile a tunebook called "Mappe". In order to do so, the wuser will create ad "Mappe" with

   * title
   * prduction note
   * list of assets
   * targetfolder
   
3. for each asset we have

   * rating - eventually a given set of ratings will be published in the "Mappe"
   * reference copy indicator
   * selection of extracts in Mappe
   * sortfield to define the sequence - this is a string which defaults to title
   
4. In order to create a Mappe, 

   * user should see a list of assets. assets might be filtered according to the metadata above, mainly by a text seard in the filename
   * user selects a "target Mappe"
   * user inserts an asset to Mappe
   * user rates the asset for the particular Mappe
   * user can preview the assed (mainly open it in Zupfnoter, even edit it in Zupfnoter)
   
5. User triggers the production of Mappe


# Architecture

## backend

1. Backend provides a grphql interface
2. Backend serves the application
3. all processing is done in the backend - we will have appropriate commandline tools for this

## User interface

1. Implemented using react

2. Approach:

    1. Menu on top
    2. Asset view
      1. list of assets with metadata
      2. sort by columns
      3. filter
      4. button to insert in mappe

    4. "mappe"
      1. Dropdown selector for mappe
      2. list of assets in Mappe with inplace setting of 
        1. rating
        2. intended Extracts
        3. sorrtfield
        
        
  






   
   
