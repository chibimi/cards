<template>
	<div class="w-100">
		<div class="row">
			<label class="col-2 col-form-label"></label>
			<div class="col-10">
				<form class="form-inline statline">
					<v-autocomplete :items="items" v-model="weapon" :get-label="getLabel" @update-items="updateItems" :component-item='template' :auto-select-one-item="false" @item-selected="inputWeapon">
					</v-autocomplete>
					<input v-model="weapon.type" type="text" class="form-control" placeholder="Type">
					<input v-model="weapon.rng" type="text" class="form-control" placeholder="RNG">
					<input v-model="weapon.pow" type="text" class="form-control" placeholder="POW">
					<input v-model="weapon.rof" type="text" class="form-control" placeholder="ROF">
					<input v-model="weapon.aoe" type="text" class="form-control" placeholder="AOE">,
				res :{{weapon}}
				res :{{search}}
				</form>
			</div>
		</div>
	</div>
</template>

<script>

import ItemTemplate from './ItemTemplate.vue'
export default {
	name: "Weapon",
	props: ["id"],
	watch: {
		id: function(newVal, oldVal) {
			console.log(newVal, oldVal);
			// watch it
			this.model = {};
			this.$http
				.get("http://localhost:9901/models/" + this.id + "/weapons")
				.then(function(res) {
					console.log(res);
					this.model = res.data;
				})
				.catch(function(err) {
					console.log(err);
				});
		}
	},
	data() {
		return {
			model: {},
			items : [],
			item: null,
			search: null,
			weapons: [
				{id:1,name:"Antiphon"},
				{id:2,name:"Rapture"},
				{id:3,name:"Rapace"},
			],
			weapon:{},
			template: ItemTemplate
		};
	},
	methods: {
		save: function() {
			this.$http
				.put("http://localhost:9901/models/"+this.id, this.model)
				.then(function(res) {
					console.log(res);
					console.log(res.data);
					if (res.status === 201){
						this.id = res.data;
					}
				});
		},
		getLabel:function(item) {
			if (!item){
				return
			}
			console.log("item", item)
      		return item.name
    	},
    	updateItems (text) {
			this.items = this.weapons.filter(item => item.name.toLowerCase().startsWith(text.toLowerCase()));
			this.search = text;
      	},
    	inputWeapon (text) {
			console.log("input",text);
			this.search = text.name
      	},
    }
	
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.statline input {
	max-width: 4rem;
}
</style>
