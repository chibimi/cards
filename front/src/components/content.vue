<template>
	<div>
        <b-form inline class="mb-3">
            <label>Faction:</label>
            <b-form-select v-model="faction.selected" :options="faction.options" />
            <label class="ml-2">Category:</label>
            <b-form-select v-model="category.selected" :options="category.options" />
            <label class="ml-2">Model:</label>
            <b-form-select v-model="model.selected" :options="model.options" />
        </b-form>

        <b-form class="mb-3">
            <b-form-input type="text" placeholder="Shortname: Fyanna2"/>
            <b-form-input type="text" placeholder="Fullname: Fyanna, Torment of Everblight"/>
            <b-form-input type="text" placeholder="Category: Blighted Strider Warlock"/>
        </b-form>

        <b-form inline class="mb-3 statline">
            <label>Statline:</label>
            <b-form-input type="text" placeholder="SPD"/>
            <b-form-input type="text" placeholder="STR"/>
            <b-form-input type="text" placeholder="MAT"/>
            <b-form-input type="text" placeholder="RAT"/>
            <b-form-input type="text" placeholder="DEF"/>
            <b-form-input type="text" placeholder="ARM"/>
            <b-form-input type="text" placeholder="CMD"/>
        </b-form>

        <b-form inline class="mb-3 statline">
            <label>Meele Weapon:</label>
            <b-form-input class="w-25" type="text" placeholder="Name"/>
            <b-form-input type="text" placeholder="Rng"/>
            <b-form-input type="text" placeholder="Pow"/>
            <b-form-input type="text" placeholder="P + S"/>
            <b-button class="ml-3" type="submit" variant="primary">Add</b-button>
        </b-form>
        

        <b-form inline class="mt-3 statline">
            <label>Spell:</label>
            <b-form-input class="w-25" type="text" placeholder="Name"/>
            <b-form-input type="text" placeholder="Cost"/>
            <b-form-input type="text" placeholder="Rng"/>
            <b-form-input type="text" placeholder="AoE"/>
            <b-form-input type="text" placeholder="Pow"/>
            <b-form-input type="text" placeholder="Dur"/>
            <b-form-input type="text" placeholder="Off"/>
            <b-button class="ml-3" type="submit" variant="primary">Add</b-button>
        </b-form>
        <b-form-textarea  placeholder="Enter something" rows="3" max-rows="6"/>
        
        <b-form inline class="mt-3">
            <label>Ability:</label>
            <b-form-input type="text" placeholder="Name"/>
            <b-button class="ml-3" type="submit" variant="primary">Add</b-button>
        </b-form>
        <b-form-textarea  placeholder="Enter something" rows="3" max-rows="6"/>
        

        <b-alert variant="success" show>Success Alert</b-alert>
    </div>
</template>

<script>
export default {
    name: "Content",
    data() {
        return {
            faction: {
                selected: null,
                options: [
                    {value:"legion", text:"Legion of Everblight"},
                    {value:"troll", text:"Trollbloods"},
                    {value:"skorne", text:"Skorne"}
                ]
            }, 
            category: {
                selected: null,
                options: [
                    {value:"warlock", text:"Warlock"},
                    {value:"warbeast", text:"Warbeast"},
                    {value:"unit", text:"Unit"},
                    {value:"solo", text:"Solo"}
                ]
            },
            model: {
                selected: null,
                options: [
                    {value:"fyanna2", text:"Fyanna, Torment of Everblight"},
                    {value:"champion", text:"BlightedNyss champion and Abbot"},
                    {value:"unit", text:"Unit"},
                    {value:"solo", text:"Solo"}
                ]
            } 
        }
    }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.statline input {
    width: 4rem;
    margin-left: 1rem;
}
</style>
