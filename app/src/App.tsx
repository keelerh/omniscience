import * as React from 'react';
import './App.css';
import './theme.css';
import {
    ActionBar,
    ActionBarRow,
    HitItemProps,
    Hits,
    HitsStats,
    Layout,
    LayoutBody,
    LayoutResults,
    NoHits,
    SearchBox,
    SearchkitManager,
    SearchkitProvider,
    TopBar,
} from 'searchkit';
import { get } from 'lodash';

// Set ES url - use a protected URL that only allows read actions.
const elasticsearchHost = 'http://localhost:9200';
const searchkit = new SearchkitManager(elasticsearchHost);

const HitItem = (props: HitItemProps) => {
    const {result, bemBlocks} = props;
    return (
        <div className={bemBlocks.item().mix(bemBlocks.container('item'))}>
            {console.log(result._source)}
            <div
                className={bemBlocks.item('title')}
                dangerouslySetInnerHTML={{__html: get(
                        result, 'highlight.title', '') || result._source.title}}
            />
            <div>
                <small
                    className={bemBlocks.item('description')}
                    dangerouslySetInnerHTML={{__html: get(result, 'highlight.description', '')}}
                />
            </div>
        </div>);
};

const App: React.SFC<{}> = () => (
    <SearchkitProvider searchkit={searchkit}>
        <Layout>
            <TopBar>
                <SearchBox
                    autofocus={true}
                    searchOnChange={false}
                    queryOptions={{analyzer: 'standard'}}
                    queryFields={['name', 'description', 'content', 'url']}
                    prefixQueryFields={['name', 'description', 'content', 'url']}
                    translations={{'NoHits.DidYouMean': 'Search for {suggestion}'}}
                />
            </TopBar>
            <LayoutBody>
                <LayoutResults>
                    <ActionBar>
                        <ActionBarRow>
                            <HitsStats/>
                        </ActionBarRow>
                    </ActionBar>
                    <Hits
                        mod="sk-hits-list"
                        hitsPerPage={15}
                        highlightFields={['title', 'description', 'url']}
                        sourceFilter={['title', 'description', 'url']}
                        itemComponent={HitItem}
                    />
                    <NoHits
                        suggestionsField={'title'}
                    />
                </LayoutResults>
            </LayoutBody>
        </Layout>
    </SearchkitProvider>
);

export default App;
