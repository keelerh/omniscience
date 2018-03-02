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

const HitItem = (props: HitItemProps) => (
    <div className={props.bemBlocks.item().mix(props.bemBlocks.container('item'))}>
        <div
            className={props.bemBlocks.item('name')}
            dangerouslySetInnerHTML={{__html: get(
                    props.result, 'highlight.name', false) || props.result._source.name}}
        />
        <div>
            <small
                className={props.bemBlocks.item('highlights')}
                dangerouslySetInnerHTML={{__html: get(props.result, 'highlight.description', '')}}
            />
        </div>
    </div>
);

const App: React.SFC<any> = () => (
    <SearchkitProvider searchkit={searchkit}>
        <Layout>
            <TopBar>
                <SearchBox
                    autofocus={true}
                    searchOnChange={false}
                    queryOptions={{analyzer: 'standard'}}
                    queryFields={['name', 'description', 'content']}
                    prefixQueryFields={['name', 'description', 'content']}
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
                        hitsPerPage={10}
                        highlightFields={['name', 'description', 'content']}
                        sourceFilter={['name', 'description', 'content']}
                        itemComponent={HitItem}
                    />
                    <NoHits
                        mod="sk-hits-list"
                        suggestionsField={'name'}
                    />
                </LayoutResults>
            </LayoutBody>
        </Layout>
    </SearchkitProvider>
);

export default App;
