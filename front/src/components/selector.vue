<template>
	<div class="row">
		<form v-on:submit.prevent class="form-inline">
			<country-flag :country="this.$language"/>
			<select :value="this.$language" class="form-control" @change="$emit('change_language', $event.target.value)">
				<option>UK</option>
				<option>FR</option>
				<option>DE</option>
				<option>IT</option>
			</select>
			<label>Faction</label>
			<select v-model="faction" class="form-control" @change="changeFaction">
				<option v-for="f in factions" :key="f.id" :value="f.id">{{f.name}}</option>
			</select>
			<label>Category</label>
			<select v-model="category" class="form-control" @change="changeCategory">
				<option v-for="c in categories" :key="c.id" :value="c.id">{{c.name}}</option>
			</select>
			<label>Ref</label>
			<select v-model="ref" class="form-control">
				<option v-for="c in refs" :key="c.id" :value="c">[{{c.status}}] #{{c.id}} {{c.title}}</option>
			</select>
			<button type="submit" class="btn btn-primary" @click="$emit('select_ref', ref)">Go</button>
			<input v-model="newName" type="text" class="form-control">
			<button type="submit" class="btn btn-primary" @click="newRef">New Ref</button>
		</form>
	</div>
</template>

<script>
import { Factions, Categories } from "./const.js";
export default {
	name: "Selector",
	components: {},
	created: function() {
		this.getRefs(this.faction, this.category);
	},
	data() {
		return {
			newName: "",
			factions: Factions,
			faction: 11,
			categories: Categories,
			category: 5,
			ref: null,
			refs: [],
		};
	},
	methods: {
		changeLanguage: function(language){
			this.$change_language(language);
		},
		getRefs: function(faction, category) {
			if (!faction || !category) {
				return;
			}
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT+ "/ref?faction_id=" + faction + "&category_id=" + category +  "&lang=" + this.$language)
				.then(function(res) {
					console.log(res);
					this.refs = res.data;
				})
				.catch(function(err) {
					console.log(err);
				});
		},
		changeFaction: function() {
			this.getRefs(this.faction, this.category);
		},
		changeCategory: function() {
			this.getRefs(this.faction, this.category);
		},
		newRef: function() {
			if (!this.faction || !this.category || !this.newName) {
				return;
			}
			var ref = {
				faction_id: this.faction,
				category_id: this.category,
				title: this.newName,
			}
			this.$http
				.post(process.env.VUE_APP_API_ENDPOINT+ "/ref?faction_id=" + this.faction + "&category_id=" + this.category, ref)
				.then(function(res) {
					console.log(res);
					this.$emit('select_ref', res.body)
					this.newName=""
				})
				.catch(function(err) {
					console.log(err);
				});
		},
	}
};
</script>

<style>
</style>
