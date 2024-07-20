import React from 'react'
import { RouteComponentProps } from 'react-router-dom';
import { Dashboard } from '../components/template/Dashboard';
import { AlertDialog, CustomizedSnackbars, OutlinedCard } from '../components/organism';
import { IItem, IIndexItemRequest, HTTPCallStatus, ICategory, IUpsertItemResponse} from '../data/interfaces';
import { serviceIndexItem } from '../data/services';
import "regenerator-runtime/runtime.js";
import { Typography, Box, TextField} from '@material-ui/core';
import { serviceAddItem, serviceEditItem } from '../data/services/ItemService';
import ItemDetailPage from './ItemDetailPage';
import { FormGroup, Label, Input } from 'reactstrap';

interface Props extends RouteComponentProps{};

interface IItemPage{
  rawContent:IItem[],
  viewConstraint:IIndexItemRequest,
  searchKey:string,
  category:string,
  categories:ICategory[]
  snackbar:{
    isShown:boolean,
    severity:string,
    msg:[]
  },
  editDialog:{
    isShown:boolean
  }
}

const getInitViewConstraint = () => ({
  owner:"",
  category:""
})


export class ItemPage extends React.Component<Props,any> {
  _isMounted:boolean = true
  state:IItemPage;
  constructor(props:Props){
    super(props);
    this.state = {
      rawContent:[],
      viewConstraint:getInitViewConstraint(),
      searchKey:"",
      category:"",
      categories:[],
      snackbar:{
        isShown:false,
        severity:"info",
        msg:[]
      },
      editDialog:{
        isShown:true
      }
    }
  }

  closeSnackbar = () => {
    this.setState({
      snackbar:{
        isShown:false,
        severity:"info",
        msg:[]
      }
    });
  }

  loadAllCategories = async () => {
    console.log("posting index cate request")
    // await serviceIndexCategory().subscribe(
    // 	(res) => {
    // 		console.log("RES:"+Object.keys(res).length);
    // 		console.log(res.data["categories"]);
    // 		this.setState({
    // 			categories: res.data["categories"]
    // 		});
    // 	},
    // 	(err)=>{
    // 		console.log("axios err:"+err);
    // 		this.setErrorSnackbar(err)
    // 	}
    // );
  }

  loadAllItems = async () => {
    console.log("posting index request")
    await serviceIndexItem(this.state.viewConstraint).subscribe(
      (res) => {
        console.log("RES:"+Object.keys(res).length);
        console.log(res.data["items"]);
        this.setState({
          rawContent: res.data["items"]
        });
        console.log("STATE:"+Object.keys(this.state.rawContent).length);
      },
      (err)=>{
        this.setErrorSnackbar(err)
      }
    );
  }

  setSuccessSnackbarEdit = (res) => {
    if(res.data['status'] == HTTPCallStatus.Success){
      // TODO: set viewConstraint to default ?
      this.loadAllItems()
    }
    this.setState({
      snackbar:{
        isShown:true,
        severity: ((res.data['status'] == HTTPCallStatus.Success) ? "success" : "error"),
        msg:res.data['message']
      }
    })
  }

  setSuccessSnackbarDelete = (res,key) => {
    console.log("deleting")
      if(res.data['status'] == HTTPCallStatus.Success){
        var array = [...this.state.rawContent]
        var index = array.map((e) => {
          return e.id
        }).indexOf(key);
        array.splice(index,1);

        this.setState({rawContent:array});
      }
      this.setState({
        snackbar:{
          isShown:true,
          severity: ((res.data['status'] == HTTPCallStatus.Success) ? "success" : "error"),
          msg:res.data['message']
        }
      })
  }

  setErrorSnackbar = (err) => {
    this.setState({
      snackbar:{
        isShown:true,
        severity:"error",
        msg:[err.message.split(),
        (err.message.includes("403") ? "please login first": "")]
      }
    })
  }

  search = (e) => {
    this.setState({
      searchKey:e.target.value
    },() => console.log(this.state.searchKey))
  }
  
  searchByCategory = (e) => {
    this.setState({
      category:e.target.value
    })
    console.log(this.state.category)
  }

  async componentDidMount(){
    this._isMounted = true;			
    if(this._isMounted){
      this.loadAllCategories()
    }
    this.loadAllItems();
  }
  componentWillUnmount(){
    this._isMounted=false;
  }
  

  render(){
    let searchKeyword: string = this.state.searchKey
    let selectedCategory : string = this.state.category
    const filteredItems = this.state.rawContent.filter(
      item => {
        if(searchKeyword === null) return 1;
        if(selectedCategory.length > 0 && item.category !== null && item.category !== selectedCategory) return 0;
        return ((item.name != null && item.name.toLowerCase().indexOf(searchKeyword.toLowerCase()) !== -1
        || (item.owner != null && item.owner.toLowerCase().indexOf(searchKeyword.toLowerCase()) !== -1)
        || (item.description != null && item.description.toLowerCase().indexOf(searchKeyword.toLowerCase()) !== -1))
        )
      }
    )

    return (
      <Dashboard 
      titlePage = {""}			
      content={
        <div>
          <div style={{display:'table'}}>
            <div style={{float: 'left', width:'25%'}}>
              <img id="loading-img" style={{width:'30%', height:'auto'}} src="https://raw.githubusercontent.com/JoshEvan/FindComputerWeb/master/findcomputer-frontend/src/assets/imgs/logo.png"/>
            </div>
            <div style={{float: 'left', width:'75%'}}>
              <Typography variant="h3" component="h2" gutterBottom>
                Solomon
              </Typography>
            </div>
          </div>
          <div>
            {
              this.state.snackbar.isShown &&
              (<CustomizedSnackbars
                severity={this.state.snackbar.severity}
                msg={this.state.snackbar.msg}
                parentCallback={this.closeSnackbar}
              />)
            }
          </div>
          <Box display="flex" flexWrap="wrap">
            <Box p={1}>
            <TextField
              id="outlined-full-width"
              label="search"
              style={{ margin: 8 }}
              placeholder="search something"
              helperText="Enter item's name / description / owner"
              fullWidth
              margin="normal"
              InputLabelProps={{
                shrink: true,
              }}
              variant="outlined"
              onChange={this.search}
            />
            </Box>
            {/* insert filter */}
          </Box>
          <Box display="flex" flexWrap="wrap">
          {
      filteredItems.map(
      (c:IItem) => {
          return(
            <React.Fragment>
              <Box p={1}>
                <OutlinedCard
                    name = {c.name}
                    price = {c.price}
                    actions={
                      <AlertDialog
                      color="primary"
                      param={c.id}
                      parentAllowance = {this.state.editDialog.isShown}
                      buttonTitle="show more"
                      parentCallbackOpen={()=>this.setState({editDialog:{isShown:true}})}
                      dialogTitle="Item Details"
                      usingAction={false}
                      dialogContent={
                        <ItemDetailPage
                          id={c.id}
                          category={c.category}
                          owner = {c.owner}
                          name = {c.name}
                          price = {c.price}
                          priceAmount = {c.priceAmount}
                          description = {c.description}
                          categories={this.state.categories}
                          parrentCallbackSuccessDelete = {this.setSuccessSnackbarDelete}
                          parrentCallbackSuccessEdit = {this.setSuccessSnackbarEdit}
                          parrentCallbackError = {this.setErrorSnackbar}
                        />
                      }
                    />
                    }
                  />
                </Box>
            </React.Fragment>
          );
        }
              )
            }
            </Box>
          </div>
        }/>
      )
  }
};